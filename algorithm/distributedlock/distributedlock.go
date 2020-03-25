package distributedlock

import (
	"fmt"
	"github.com/CodisLabs/codis/pkg/utils/redis"
	"sync"
	"time"
)

type redisConfig struct {
	key    string
	ip     string
	port   string
	expire time.Duration
}

/* 互斥锁会降低处理性能，考虑有没有更好的方式 */
type DistributedLock struct {
	config *redisConfig
	client *redis.Client
	mutex  sync.Mutex
}

func NewDistributedLock(key string, ip string, port string, expire time.Duration) (*DistributedLock, error) {
	d := &DistributedLock{
		config:&redisConfig{
			key:key,
			ip:ip,
			port:port,
			expire:expire,
		},
	}
	c, err := redis.NewClient(ip+":"+port, "", expire)
	if err == nil {
		d.client = c
	}

	return d, nil
}

/* 延续锁的生命周期，延长的时间为设置的超时时长 */
func (d *DistributedLock)ContinueLife(value string) (bool, error) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	cmd := fmt.Sprintf("get %s", d.config.key)
	r, err := d.client.Do(cmd)
	if err == nil && r == value {
		cmd := fmt.Sprintf("expire %s %d", d.config.key, d.config.expire)
		if _, err := d.client.Do(cmd); err != nil {
			return false, fmt.Errorf("continue life failed err: %s", err)
		}
	} else {
		if err == nil {
			return false, fmt.Errorf("not you")
		}
	}
	return true, nil
}

/* 加分布式锁，增加超时时间的目的是防止上锁的节点异常导致锁不会被解开
 不同的访问者可以设置不同的value，用于区分彼此。后续可以增加解锁时对解锁方的验证。
*/
func (d *DistributedLock)Lock(value string) (bool, error) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	cmd := fmt.Sprintf("set %s %s %d NX", d.config.key, value, d.config.expire)
	_, err := d.client.Do(cmd)
	if err != nil {
		return false, fmt.Errorf("lock failed err: %s", err)
	}

	return true, nil
}

/* 加互斥锁保证修改value的操作原子性，或者使用redis事务。
   先从redis中取出value，验证解锁的是不是加锁的。
*/
func (d *DistributedLock)UnLock(value string) (bool, error) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	cmd := fmt.Sprintf("get %s", d.config.key)
	r, err := d.client.Do(cmd)
	if r == value && err == nil {
		cmd := fmt.Sprintf("delete %s", d.config.key)
		_, err := d.client.Do(cmd)
		if err != nil {
			return false, fmt.Errorf("delete key failed err: %s", err)
		}
	} else if err != nil {
		return false, fmt.Errorf("get key failed err: %s", err)
	} else {
		return false, fmt.Errorf("not you")
	}

	return true, nil
}