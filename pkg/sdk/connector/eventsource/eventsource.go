package eventsource

import (
	"context"
	"sync"

	"github.com/kuro-jojo/go-utils/pkg/api/models"
	"github.com/kuro-jojo/go-utils/pkg/sdk/connector/types"
)

// EventSource is anything that can be used
// to get events from the Keptn Control Plane
type EventSource interface {
	// Start triggers the execution of the EventSource
	Start(context.Context, types.RegistrationData, chan types.EventUpdate, chan error, *sync.WaitGroup) error
	// OnSubscriptionUpdate can be called to tell the EventSource that
	// the current subscriptions have been changed
	OnSubscriptionUpdate([]models.EventSubscription)
	// Sender returns a component that gives the possiblity to send events back
	// to the Keptn Control plane
	Sender() types.EventSender
	//Stop is stopping the EventSource
	Stop() error
	// Cleanup performs all cleanup actions that should be executed before the component gets shut down
	Cleanup() error
}
