package main

import (
  "testing";"time"
)

func TestAfter(t *testing.T){

	startTime:=time.Now();
	sleep();
	endTime:=time.Since(startTime);

	if int64(endTime)/1000000000!=int64(time.Second * time.Duration(5))/1000000000{
		t.Error("Sleep is not working")
	}
}
