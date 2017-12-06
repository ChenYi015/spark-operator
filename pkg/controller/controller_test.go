package controller

import (
	"testing"

	"github.com/liyinan926/spark-operator/pkg/apis/v1alpha1"
	"github.com/liyinan926/spark-operator/pkg/crd"
	"github.com/stretchr/testify/assert"

	apiv1 "k8s.io/api/core/v1"
	kubeclientfake "k8s.io/client-go/kubernetes/fake"
	apiextensionsfake "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/fake"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestOnAdd(t *testing.T) {
}

func newFakeController() *SparkApplicationController {
	crdClient := crd.NewFakeClient()
	kubeClient := kubeclientfake.NewSimpleClientset()
	kubeClient.CoreV1().Nodes().Create(&apiv1.Node{
		ObjectMeta: metav1.ObjectMeta{
			Name: "node1",
		},
		Status: apiv1.NodeStatus{
			Addresses: []apiv1.NodeAddress{
				{
					Type: apiv1.NodeExternalIP,
					Address: "12.34.56.78",
				},
			},
		},
	})
	return New(crdClient, kubeClient, apiextensionsfake.NewSimpleClientset(), 1)
}

func TestProcessSingleDriverStateUpdate(t *testing.T) {
	type testcase struct {
		name string
		update driverStateUpdate
		expectedAppState v1alpha1.ApplicationStateType
	}

	ctrl := newFakeController()

	appID := "foo-123"
	app := &v1alpha1.SparkApplication{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foo",
			Namespace: "default",
		},
		Status: v1alpha1.SparkApplicationStatus{
			AppID: appID,
			AppState: v1alpha1.ApplicationState{
				State:        v1alpha1.NewState,
				ErrorMessage: "",
			},
		},
	}
	ctrl.runningApps[app.Status.AppID] = app

	testcases := []testcase{
		{
			name: "succeeded driver",
			update: driverStateUpdate{
				appID:    appID,
				podName:  "foo-driver",
				nodeName: "node1",
				podPhase: apiv1.PodSucceeded,
			},
			expectedAppState: v1alpha1.CompletedState,
		},
		{
			name: "failed driver",
			update: driverStateUpdate{
				appID:    appID,
				podName:  "foo-driver",
				nodeName: "node1",
				podPhase: apiv1.PodFailed,
			},
			expectedAppState: v1alpha1.FailedState,
		},
		{
			name: "running driver",
			update: driverStateUpdate{
				appID:    appID,
				podName:  "foo-driver",
				nodeName: "node1",
				podPhase: apiv1.PodRunning,
			},
			expectedAppState: v1alpha1.NewState,
		},
	}

	testFn := func(test testcase, t *testing.T) {
		ctrl.runningApps[test.update.appID].Status.AppState.State = v1alpha1.NewState
		ctrl.processSingleDriverStateUpdate(test.update)
		app, ok := ctrl.runningApps[test.update.appID]
		if !ok {
			t.Errorf("%s: SparkApplication %s not found", test.name, test.update.appID)
		}
		assert.Equal(
			t,
			app.Status.AppState.State,
			test.expectedAppState,
			"wanted application state %s got %s",
			test.expectedAppState,
			app.Status.AppState.State)
	}

	for _, test := range testcases {
		testFn(test, t)
	}
}

func TestProcessSingleAppStateUpdate(t *testing.T) {
	type testcase struct {
		name string
		update appStateUpdate
		initialAppState v1alpha1.ApplicationStateType
		expectedAppState v1alpha1.ApplicationStateType
	}

	ctrl := newFakeController()

	appID := "foo-123"
	app := &v1alpha1.SparkApplication{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foo",
			Namespace: "default",
		},
		Status: v1alpha1.SparkApplicationStatus{
			AppID: appID,
			AppState: v1alpha1.ApplicationState{
				State:        v1alpha1.NewState,
				ErrorMessage: "",
			},
		},
	}
	ctrl.runningApps[app.Status.AppID] = app

	testcases := []testcase{
		{
			name: "succeeded app",
			update: appStateUpdate{
				appID:    appID,
				state:  v1alpha1.CompletedState,
				errorMessage: "",
			},
			initialAppState: v1alpha1.RunningState,
			expectedAppState: v1alpha1.CompletedState,
		},
		{
			name: "failed app",
			update: appStateUpdate{
				appID:    appID,
				state:  v1alpha1.FailedState,
				errorMessage: "",
			},
			initialAppState: v1alpha1.RunningState,
			expectedAppState: v1alpha1.FailedState,
		},
		{
			name: "running app",
			update: appStateUpdate{
				appID:    appID,
				state:  v1alpha1.RunningState,
				errorMessage: "",
			},
			initialAppState: v1alpha1.NewState,
			expectedAppState: v1alpha1.RunningState,
		},
		{
			name: "completed app with initial state SubmittedState",
			update: appStateUpdate{
				appID:    appID,
				state:  v1alpha1.SubmittedState,
				errorMessage: "",
			},
			initialAppState: v1alpha1.CompletedState,
			expectedAppState: v1alpha1.CompletedState,
		},
		{
			name: "failed app with initial state SubmittedState",
			update: appStateUpdate{
				appID:    appID,
				state:  v1alpha1.SubmittedState,
				errorMessage: "",
			},
			initialAppState: v1alpha1.FailedState,
			expectedAppState: v1alpha1.FailedState,
		},
	}

	testFn := func(test testcase, t *testing.T) {
		ctrl.runningApps[test.update.appID].Status.AppState.State = test.initialAppState
		ctrl.processSingleAppStateUpdate(test.update)
		app, ok := ctrl.runningApps[test.update.appID]
		if !ok {
			t.Errorf("%s: SparkApplication %s not found", test.name, test.update.appID)
		}
		assert.Equal(
			t,
			app.Status.AppState.State,
			test.expectedAppState,
			"wanted application state %s got %s",
			test.expectedAppState,
			app.Status.AppState.State)
	}

	for _, test := range testcases {
		testFn(test, t)
	}
}

func TestProcessSingleExecutorStateUpdate(t *testing.T) {
	type testcase struct {
		name string
		update executorStateUpdate
		expectedExecutorStates map[string]v1alpha1.ExecutorState
	}

	ctrl := newFakeController()

	appID := "foo-123"
	app := &v1alpha1.SparkApplication{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foo",
			Namespace: "default",
		},
		Status: v1alpha1.SparkApplicationStatus{
			AppID: appID,
			AppState: v1alpha1.ApplicationState{
				State:        v1alpha1.NewState,
				ErrorMessage: "",
			},
			ExecutorState: make(map[string]v1alpha1.ExecutorState),
		},
	}
	ctrl.runningApps[app.Status.AppID] = app

	testcases := []testcase{
		{
			name: "completed executor",
			update: executorStateUpdate{
				appID: appID,
				podName: "foo-exec-1",
				executorID: "1",
				state: v1alpha1.ExecutorCompletedState,
			},
			expectedExecutorStates: map[string]v1alpha1.ExecutorState {
				"foo-exec-1": v1alpha1.ExecutorCompletedState,
			},
		},
		{
			name: "failed executor",
			update: executorStateUpdate{
				appID: appID,
				podName: "foo-exec-2",
				executorID: "2",
				state: v1alpha1.ExecutorFailedState,
			},
			expectedExecutorStates: map[string]v1alpha1.ExecutorState {
				"foo-exec-1": v1alpha1.ExecutorCompletedState,
				"foo-exec-2": v1alpha1.ExecutorFailedState,
			},
		},
		{
			name: "running executor",
			update: executorStateUpdate{
				appID: appID,
				podName: "foo-exec-3",
				executorID: "3",
				state: v1alpha1.ExecutorRunningState,
			},
			expectedExecutorStates: map[string]v1alpha1.ExecutorState {
				"foo-exec-1": v1alpha1.ExecutorCompletedState,
				"foo-exec-2": v1alpha1.ExecutorFailedState,
				"foo-exec-3": v1alpha1.ExecutorRunningState,
			},
		},
		{
			name: "pending executor",
			update: executorStateUpdate{
				appID: appID,
				podName: "foo-exec-4",
				executorID: "4",
				state: v1alpha1.ExecutorPendingState,
			},
			expectedExecutorStates: map[string]v1alpha1.ExecutorState {
				"foo-exec-1": v1alpha1.ExecutorCompletedState,
				"foo-exec-2": v1alpha1.ExecutorFailedState,
				"foo-exec-3": v1alpha1.ExecutorRunningState,
			},
		},
	}

	testFn := func(test testcase, t *testing.T) {
		ctrl.processSingleExecutorStateUpdate(test.update)
		app, ok := ctrl.runningApps[test.update.appID]
		if !ok {
			t.Errorf("%s: SparkApplication %s not found", test.name, test.update.appID)
		}
		assert.Equal(
			t,
			app.Status.ExecutorState,
			test.expectedExecutorStates,
			"wanted executor states %s got %s",
			test.expectedExecutorStates,
			app.Status.ExecutorState)
	}

	for _, test := range testcases {
		testFn(test, t)
	}
}