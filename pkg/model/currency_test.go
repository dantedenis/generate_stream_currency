package model

import (
	"github.com/stretchr/testify/assert"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"
)

func Test_Rate(t *testing.T) {
	c := NewRate()
	var tm time.Time
	for i := 0; i < 150; i++ {
		tm = time.Unix(rand.Int63(), 0)
		c.Add(tm, 123123.13)
	}

}

func Test(t *testing.T) {
	timer := time.Now()
	log.Println(timer.Unix())
}

func TestCurrency_RunGenerate(t *testing.T) {
	c := &Currency{}

	assert.Nil(t, c.RunGenerate())
	assert.NotNil(t, c.GetVal("USDRUB"))
}

func TestCurrency_RunGenerate_error1(t *testing.T) {
	err := os.Setenv("CHANCE", "sdfdf")
	if err != nil {
		t.Error(err)
		return
	}
	c := &Currency{}
	assert.NotNil(t, c.RunGenerate())

}

func TestCurrency_RunGenerate_error2(t *testing.T) {
	err := os.Setenv("END_TIME", "sdfdf")
	if err != nil {
		t.Error(err)
		return
	}
	c := &Currency{}
	assert.NotNil(t, c.RunGenerate())

}

func TestCurrency_RunGenerate_error3(t *testing.T) {
	err := os.Setenv("START_TIME", "vbxb")
	if err != nil {
		t.Error(err)
		return
	}
	c := &Currency{}
	assert.NotNil(t, c.RunGenerate())

}

func TestCurrency_RunGenerate_error_start_on_end(t *testing.T) {
	start := os.Getenv("START_TIME")
	end := os.Getenv("END_TIME")

	assert.Nil(t, os.Setenv("START_TIME", end))
	assert.Nil(t, os.Setenv("END_TIME", start))

	c := &Currency{}
	assert.NotNil(t, c.RunGenerate())

}
