package lock

import (
	// "log"

	

	"6.5840/kvsrv1/rpc"
	"6.5840/kvtest1"
)

type Lock struct {
	// IKVClerk is a go interface for k/v clerks: the interface hides
	// the specific Clerk type of ck but promises that ck supports
	// Put and Get.  The tester passes the clerk in when calling
	// MakeLock().
	ck kvtest.IKVClerk
	ClientID string
    LockName string 
	// You may add code here
}

// The tester calls MakeLock() and passes in a k/v clerk; your code can
// perform a Put or Get by calling lk.ck.Put() or lk.ck.Get().
//
// Use l as the key to store the "lock state" (you would have to decide
// precisely what the lock state is).
func MakeLock(ck kvtest.IKVClerk, l string) *Lock {
	lk := &Lock{ck: ck}
	lk.LockName=l
	lk.ClientID=kvtest.RandValue(8)
	return lk
}

func (lk *Lock) Acquire() {
	// Your code here
	for{
		val,ver,err:=lk.ck.Get(lk.LockName)
		//没有对应的锁
		if err==rpc.ErrNoKey{
			ok:=lk.ck.Put(lk.LockName,lk.ClientID,0)
		    if ok==rpc.OK||ok==rpc.ErrMaybe{
				break
			}
		}else if err==rpc.OK{
			//有对应的锁
			if val==""{
				ok:=lk.ck.Put(lk.LockName,lk.ClientID,ver)
				if ok==rpc.OK{
					break
				}
			}else if val!=lk.ClientID{
				continue
			}else if val==lk.ClientID{
				break
			}
		}		
	}
	
}

func (lk *Lock) Release() {
	// Your code here
	val,ver,err:=lk.ck.Get(lk.LockName)
	// log.Printf("Release:val=%v,ver=%v,err=%v,lk.ClientID:%v",val,ver,err,lk.ClientID)
	if err==rpc.OK{
		if val==lk.ClientID{
			lk.ck.Put(lk.LockName,"",ver)
		}
	}
}
