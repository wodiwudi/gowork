package ch8

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

//对象池单例
type ReusableObj struct {
}

//对象池
type ObjPool struct {
	bufchan chan *ReusableObj
}

//初始化对象池
func NewObjPool(nums int) *ObjPool {
	objPool := ObjPool{}
	objPool.bufchan = make(chan *ReusableObj, nums)
	for i := 0; i < nums; i++ {
		objPool.bufchan <- &ReusableObj{}
	}
	return &objPool
}

//获取一个对象
func (p *ObjPool) GetObject(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret := <-p.bufchan:
		return ret, nil
	case <-time.After(timeout):
		return nil, errors.New("time out")
	}
}

//放回资源到资源池
func (p *ObjPool) ReleaseObject(obj *ReusableObj) error {
	select {
	case p.bufchan <- obj:
		return nil
		//如果放满了自动到default
	default:
		return errors.New("overflow")
	}
}

func TestObjPool(t *testing.T) {
	pool := NewObjPool(10)
	if err := pool.ReleaseObject(&ReusableObj{}); err != nil {
		t.Error(err)
	}
	for i := 0; i < 11; i++ {
		if v, err := pool.GetObject(time.Second * 1); err != nil {
			t.Error(err)
		} else {
			fmt.Printf("%T\n", v)
			/*err := pool.ReleaseObject(v)
			if err != nil {
				t.Error(err)
			}*/
		}
	}
	fmt.Print("Done")
}
