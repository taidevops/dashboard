package sync

import (
	"sync"
	"time"

	"k8s.io/apimachinery/pkg/watch"
	v1 "k8s.io/client-go/informers/coordination/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/watch"
)

const secretSyncPeriod = 5 * time.Minute

type secretSynchonizer struct {
	namespace string
	name string

	secret *v1.Secret
	client kubernetes.Interface
	actionHandlers map[watch.EventType][]syncApi.ActionHandlerFunction
	errChan chan error
	poller syncApi.Poller

	mux sync.Mutex
}

