package ch8

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

type ResObj struct {

}

type ResObjPool struct {
	bufChan chan *ResObj
}

func NewObjPool(numOfObj int) *ResObjPool {
	pool := ResObjPool{bufChan: make(chan *ResObj, numOfObj)}
	for i := 0;i < numOfObj;i++{
		pool.bufChan <- &ResObj{}
	}
	return &pool
}
func (p *ResObjPool) GetOneObj()(*ResObj,error)  {
	select {
	case res := <-p.bufChan:
		return res,nil
	case  <- time.After(time.Millisecond * 100):
		return nil,errors.New("获取对象池超时")
	}
}

func (p *ResObjPool)putIn(obj *ResObj) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}

func TestObhPool(t *testing.T) {
	pool := NewObjPool(10)

	//err := pool.putIn(&ResObj{})// 只有10个 不然会出错
	//fmt.Println(err)
	for{
		obj, err := pool.GetOneObj()
		if err == nil {
			fmt.Println(obj)
			err := pool.putIn(obj)
			if err != nil{
				fmt.Println(err)
			}
		}else {
			// 否则break
			fmt.Println(err)
			break
		}


	}
}