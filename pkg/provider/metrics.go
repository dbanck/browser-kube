package provider

import (
	"context"

	stats "k8s.io/kubernetes/pkg/kubelet/apis/stats/v1alpha1"
)

// GetStatsSummary returns the stats summary for pods running on ACI
func (p *BrowserProvider) GetStatsSummary(ctx context.Context) (summary *stats.Summary, err error) {
	return nil, nil
	// ctx, span := trace.StartSpan(ctx, "GetSummaryStats")
	// defer span.End()

	// p.metricsSync.Lock()
	// defer p.metricsSync.Unlock()

	// log.G(ctx).Debug("acquired metrics mutex")

	// if time.Now().Sub(p.metricsSyncTime) < time.Minute {
	// 	span.WithFields(ctx, log.Fields{
	// 		"preCachedResult":        true,
	// 		"cachedResultSampleTime": p.metricsSyncTime.String(),
	// 	})
	// 	return p.lastMetric, nil
	// }
	// ctx = span.WithFields(ctx, log.Fields{
	// 	"preCachedResult":        false,
	// 	"cachedResultSampleTime": p.metricsSyncTime.String(),
	// })

	// select {
	// case <-ctx.Done():
	// 	return nil, ctx.Err()
	// default:
	// }

	// defer func() {
	// 	if err != nil {
	// 		return
	// 	}
	// 	p.lastMetric = summary
	// 	p.metricsSyncTime = time.Now()
	// }()

	// pods, err := p.GetPods(ctx)

	// if err != nil {
	// 	return nil, errors.Wrap(err, "Getting Pods for Stats Summary failed")
	// }

	// var s stats.Summary
	// s.Node = stats.NodeStats{
	// 	NodeName: p.nodeName,
	// }
	// s.Pods = make([]stats.PodStats, 0, len(pods))
	// for _, pod := range pods {
	// 	s.Pods = append(s.Pods, *p.GetPodStats(ctx, pod.Namespace, pod.Name))
	// }

	// return &s, nil
}
