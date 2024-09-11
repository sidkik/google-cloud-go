// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

//go:build go1.23

package aiplatform

import (
	"iter"

	aiplatformpb "cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	longrunningpb "cloud.google.com/go/longrunning/autogen/longrunningpb"
	"github.com/googleapis/gax-go/v2/iterator"
	locationpb "google.golang.org/genproto/googleapis/cloud/location"
)

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *AnnotationIterator) All() iter.Seq2[*aiplatformpb.Annotation, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ArtifactIterator) All() iter.Seq2[*aiplatformpb.Artifact, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *BatchPredictionJobIterator) All() iter.Seq2[*aiplatformpb.BatchPredictionJob, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ContextIterator) All() iter.Seq2[*aiplatformpb.Context, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *CustomJobIterator) All() iter.Seq2[*aiplatformpb.CustomJob, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *DataItemIterator) All() iter.Seq2[*aiplatformpb.DataItem, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *DataItemViewIterator) All() iter.Seq2[*aiplatformpb.DataItemView, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *DataLabelingJobIterator) All() iter.Seq2[*aiplatformpb.DataLabelingJob, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *DatasetIterator) All() iter.Seq2[*aiplatformpb.Dataset, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *DatasetVersionIterator) All() iter.Seq2[*aiplatformpb.DatasetVersion, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *DeployedModelIterator) All() iter.Seq2[*aiplatformpb.DeployedModel, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *DeploymentResourcePoolIterator) All() iter.Seq2[*aiplatformpb.DeploymentResourcePool, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *EndpointIterator) All() iter.Seq2[*aiplatformpb.Endpoint, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *EntityTypeIterator) All() iter.Seq2[*aiplatformpb.EntityType, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ExecutionIterator) All() iter.Seq2[*aiplatformpb.Execution, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *FeatureGroupIterator) All() iter.Seq2[*aiplatformpb.FeatureGroup, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *FeatureIterator) All() iter.Seq2[*aiplatformpb.Feature, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *FeatureOnlineStoreIterator) All() iter.Seq2[*aiplatformpb.FeatureOnlineStore, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *FeatureViewIterator) All() iter.Seq2[*aiplatformpb.FeatureView, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *FeatureViewSyncIterator) All() iter.Seq2[*aiplatformpb.FeatureViewSync, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *FeaturestoreIterator) All() iter.Seq2[*aiplatformpb.Featurestore, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *HyperparameterTuningJobIterator) All() iter.Seq2[*aiplatformpb.HyperparameterTuningJob, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *IndexEndpointIterator) All() iter.Seq2[*aiplatformpb.IndexEndpoint, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *IndexIterator) All() iter.Seq2[*aiplatformpb.Index, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *LocationIterator) All() iter.Seq2[*locationpb.Location, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *MetadataSchemaIterator) All() iter.Seq2[*aiplatformpb.MetadataSchema, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *MetadataStoreIterator) All() iter.Seq2[*aiplatformpb.MetadataStore, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *MigratableResourceIterator) All() iter.Seq2[*aiplatformpb.MigratableResource, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ModelDeploymentMonitoringJobIterator) All() iter.Seq2[*aiplatformpb.ModelDeploymentMonitoringJob, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ModelEvaluationIterator) All() iter.Seq2[*aiplatformpb.ModelEvaluation, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ModelEvaluationSliceIterator) All() iter.Seq2[*aiplatformpb.ModelEvaluationSlice, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ModelIterator) All() iter.Seq2[*aiplatformpb.Model, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ModelMonitoringStatsAnomaliesIterator) All() iter.Seq2[*aiplatformpb.ModelMonitoringStatsAnomalies, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *NasJobIterator) All() iter.Seq2[*aiplatformpb.NasJob, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *NasTrialDetailIterator) All() iter.Seq2[*aiplatformpb.NasTrialDetail, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *NotebookExecutionJobIterator) All() iter.Seq2[*aiplatformpb.NotebookExecutionJob, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *NotebookRuntimeIterator) All() iter.Seq2[*aiplatformpb.NotebookRuntime, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *NotebookRuntimeTemplateIterator) All() iter.Seq2[*aiplatformpb.NotebookRuntimeTemplate, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *OperationIterator) All() iter.Seq2[*longrunningpb.Operation, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *PersistentResourceIterator) All() iter.Seq2[*aiplatformpb.PersistentResource, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *PipelineJobIterator) All() iter.Seq2[*aiplatformpb.PipelineJob, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *SavedQueryIterator) All() iter.Seq2[*aiplatformpb.SavedQuery, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *ScheduleIterator) All() iter.Seq2[*aiplatformpb.Schedule, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *SpecialistPoolIterator) All() iter.Seq2[*aiplatformpb.SpecialistPool, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *StudyIterator) All() iter.Seq2[*aiplatformpb.Study, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *TensorboardExperimentIterator) All() iter.Seq2[*aiplatformpb.TensorboardExperiment, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *TensorboardIterator) All() iter.Seq2[*aiplatformpb.Tensorboard, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *TensorboardRunIterator) All() iter.Seq2[*aiplatformpb.TensorboardRun, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *TensorboardTimeSeriesIterator) All() iter.Seq2[*aiplatformpb.TensorboardTimeSeries, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *TimeSeriesDataPointIterator) All() iter.Seq2[*aiplatformpb.TimeSeriesDataPoint, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *TrainingPipelineIterator) All() iter.Seq2[*aiplatformpb.TrainingPipeline, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *TrialIterator) All() iter.Seq2[*aiplatformpb.Trial, error] {
	return iterator.RangeAdapter(it.Next)
}

// All returns an iterator. If an error is returned by the iterator, the
// iterator will stop after that iteration.
func (it *TuningJobIterator) All() iter.Seq2[*aiplatformpb.TuningJob, error] {
	return iterator.RangeAdapter(it.Next)
}
