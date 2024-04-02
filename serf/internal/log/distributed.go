package log

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/hashicorp/raft"
	raftboltdb "github.com/hashicorp/raft-boltdb"
	api "github.com/travisjeffery/proglog/api/v1"
	"google.golang.org/protobuf/proto"
	"io"
	"net"
	"os"
	"path/filepath"
	"time"
)

type DistributedLog struct {
	config Config
	raft   *raft.Raft
	log    *Log
}

func NewDistributedLog(dataDir string, config Config) (
	*DisttributedLog, error) {
	l := &DistributedLog{
		config: config,
	}
	if err := l.setupLog(dataDir); err != nil {
		return nil, err
	}
	if err := l.setupRaft(); err != nil {
		return nil, err
	}
	return l, nil

}

func (l *DistributedLog) setupLog(dataDir string) error {
	logDir := filepath.Join(dataDir, "log")
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return err
	}
	l.log, err = NewLog(logDir, l.config)
	return err
}

func (l *DistributedLog) setupRaft(dataDir:string) error {

	fsm := &fsm{log: l.log}
	logDir := filepath.Join(dataDir, "raft")
	if err := os.MkdirAll(logDir, 0755); err != nil {
		return err
	}
	logCo
}