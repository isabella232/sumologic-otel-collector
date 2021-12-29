# Sumo Logic Distribution of OpenTelemetry

[![Default branch build](https://github.com/SumoLogic/sumologic-otel-collector/actions/workflows/dev_builds.yml/badge.svg)](https://github.com/SumoLogic/sumologic-otel-collector/actions/workflows/dev_builds.yml)

Sumo Logic Distro of [OpenTelemetry Collector][otc_link] built with
[opentelemetry-collector-builder][otc_builder_link].

[otc_link]: https://github.com/open-telemetry/opentelemetry-collector
[otc_builder_link]: https://github.com/open-telemetry/opentelemetry-collector-builder

**This software is currently in beta and is not recommended for production environments.**
**If you wish to participate in this beta, please contact your Sumo Logic account team or Sumo Logic Support.**

- [Usage and configuration examples](#usage-and-configuration-examples)
- [Migration from Sumo Logic Installed Collector](#migration-from-sumo-logic-installed-collector)
- [Open Telemetry collector builder](#open-telemetry-collector-builder)
- [Built-in Components](#built-in-components)
  - [Receivers](#receivers)
  - [Processors](#processors)
  - [Exporters](#exporters)
  - [Extensions](#extensions)
- [Contributing](#contributing)

## Usage and configuration examples

See the [documentation](docs/README.md).

## Migration from Sumo Logic Installed Collector

See the [migration documentation](docs/Migration.md).

## Open Telemetry collector builder

Sumo Logic Distribution of OpenTelemetry uses
[`opentelemetry-collector-builder`][otcbuilder] in order to produce the collector
binary.

One can find the details in [here][otcbuilder_dir].

[otcbuilder]: https://github.com/open-telemetry/opentelemetry-collector-builder
[otcbuilder_dir]: ./otelcolbuilder/README.md

## Built-in Components

This sections represents the supported components that are included in Sumo Logic
OT distro.

<!-- markdownlint-disable MD013 -->
<!-- markdownlint-disable MD034 -->

### Receivers

#### Sumo Logic supported receivers

| Name                                                           | Source                                                                                        |
|----------------------------------------------------------------|-----------------------------------------------------------------------------------------------|
| `telegrafreceiver` [configuration help][telegrafreceiver_help] | https://github.com/SumoLogic/sumologic-otel-collector/tree/main/pkg/receiver/telegrafreceiver |

[telegrafreceiver_help]: ./docs/Configuration.md#telegraf-receiver

#### Upstream receivers

| Name                                                                     | Source                                                                                                                 |
|--------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------|
| `awscontainerinsightreceiver`                                            | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/receiver/awscontainerinsightreceiver    |
| `awsecscontainermetricsreceiver`                                         | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/receiver/awsecscontainermetricsreceiver |
| `awsxrayreceiver`                                                        | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/receiver/awsxrayreceiver                |
| `carbonreceiver`                                                         | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/receiver/carbonreceiver                 |
| `collectdreceiver`                                                       | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/receiver/collectdreceiver               |
| `dockerstatsreceiver`                                                    | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/receiver/dockerstatsreceiver            |
| `dotnetdiagnosticsreceiver`                                              | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/receiver/dotnetdiagnosticsreceiver      |
| `filelogreceiver` [configuration help][filelogreceiver_help]             | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/receiver/filelogreceiver                |
| `fluentforwardreceiver` [configuration help][fluentforwardreceiver_help] | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/receiver/fluentforwardreceiver          |
| `googlecloudspannerreceiver`                                             | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/receiver/googlecloudspannerreceiver     |
| `hostmetricsreceiver` [configuration help][hostmetricsreceiver_help]     | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/receiver/hostmetricsreceiver            |
| `jaegerreceiver` [configuration help][jaegerreceiver_help]               | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/receiver/jaegerreceiver                 |
| `jmxreceiver`                                                            | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/receiver/jmxreceiver                    |
| `journaldreceiver`                                                       | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/receiver/journaldreceiver               |
| `kafkametricsreceiver`                                                   | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/receiver/kafkametricsreceiver           |
| `kafkareceiver`                                                          | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/receiver/kafkareceiver                  |
| `opencensusreceiver` [configuration help][opencensusreceiver_help]       | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/receiver/opencensusreceiver             |
| `podmanreceiver`                                                         | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/receiver/podmanreceiver                 |
| `receivercreator`                                                        | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/receiver/receivercreator                |
| `redisreceiver`                                                          | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/receiver/redisreceiver                  |
| `sapmreceiver`                                                           | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/receiver/sapmreceiver                   |
| `signalfxreceiver`                                                       | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/receiver/signalfxreceiver               |
| `splunkhecreceiver`                                                      | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/receiver/splunkhecreceiver              |
| `syslogreceiver` [configuration help][syslogreceiver_help]               | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/receiver/syslogreceiver                 |
| `statsdreceiver` [configuration help][statsdreceiver_help]               | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/receiver/statsdreceiver                 |
| `tcplogreceiver` [configuration help][tcplogreceiver_help]               | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/receiver/tcplogreceiver                 |
| `udplogreceiver` [configuration help][udplogreceiver_help]               | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/receiver/udplogreceiver                 |
| `wavefrontreceiver`                                                      | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/receiver/wavefrontreceiver              |
| `windowsperfcountersreceiver`                                            | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/receiver/windowsperfcountersreceiver    |
| `zipkinreceiver` [configuration help][zipkinreceiver_help]               | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/receiver/zipkinreceiver                 |
| `zookeeperreceiver`                                                      | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/receiver/zookeeperreceiver              |

[filelogreceiver_help]: ./docs/Configuration.md#filelog-receiver
[fluentforwardreceiver_help]: ./docs/Configuration.md#fluent-forward-receiver
[hostmetricsreceiver_help]: ./docs/Configuration.md#host-metrics-receiver
[jaegerreceiver_help]: ./docs/Configuration.md#jaeger-receiver
[opencensusreceiver_help]: ./docs/Configuration.md#opencensus-receiver
[statsdreceiver_help]: ./docs/Configuration.md#statsd-receiver
[syslogreceiver_help]: ./docs/Configuration.md#syslog-receiver
[tcplogreceiver_help]: ./docs/Configuration.md#tcplog-receiver
[udplogreceiver_help]: ./docs/Configuration.md#udplog-receiver
[zipkinreceiver_help]: ./docs/Configuration.md#zipkin-receiver

### Processors

#### Sumo Logic supported processors

| Name                                                                           | Source                                                                                                 |
|--------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------|
| `cascadingfilterprocessor` [configuration help][cascadingfilterprocessor_help] | https://github.com/SumoLogic/sumologic-otel-collector/tree/main/pkg/processor/cascadingfilterprocessor |
| `k8sprocessor` [configuration help][k8sprocessor_help]                         | https://github.com/SumoLogic/sumologic-otel-collector/tree/main/pkg/processor/k8sprocessor             |
| `metricfrequencyprocessor` [configuration_help][metricfrequencyprocessor_help] | https://github.com/SumoLogic/sumologic-otel-collector/tree/main/pkg/processor/metricfrequencyprocessor |
| `sourceprocessor` [configuration help][sourceprocessor_help]                   | https://github.com/SumoLogic/sumologic-otel-collector/tree/main/pkg/processor/sourceprocessor          |
| `sumologicsyslogprocessor` [configuration help][sumologicsyslogprocessor_help] | https://github.com/SumoLogic/sumologic-otel-collector/tree/main/pkg/processor/sumologicsyslogprocessor |

[cascadingfilterprocessor_help]: ./docs/Configuration.md#cascading-filter-processor
[k8sprocessor_help]: ./docs/Configuration.md#kubernetes-processor
[metricfrequencyprocessor_help]: ./docs/Configuration.md#metric-frequency-processor
[sourceprocessor_help]: ./docs/Configuration.md#source-processor
[sumologicsyslogprocessor_help]: ./docs/Configuration.md#sumo-logic-syslog-processor

#### Upstream processors

| Name                                                                               | Source                                                                                                                 |
|------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------|
| `attributesprocessor` [configuration help][attributesprocessor_help]               | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/processor/attributesprocessor           |
| `filterprocessor` [configuration help][filterprocessor_help]                       | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/processor/filterprocessor               |
| `groupbyattrsprocessor` [configuration help][groupbyattrsprocessor_help]           | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/processor/groupbyattrsprocessor         |
| `groupbytraceprocessor` [configuration help][groupbytraceprocessor_help]           | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/processor/groupbytraceprocessor         |
| `metricstransformprocessor` [configuration help][metricstransformprocessor_help]   | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/processor/metricstransformprocessor     |
| `probabilisticsamplerprocessor`                                                    | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/processor/probabilisticsamplerprocessor |
| `resourcedetectionprocessor` [configuration help][resourcedetectionprocessor_help] | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/processor/resourcedetectionprocessor    |
| `resourceprocessor` [configuration help][resourceprocessor_help]                   | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/processor/resourceprocessor             |
| `routingprocessor` [configuration help][routingprocessor_help]                     | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/processor/routingprocessor              |
| `spanmetricsprocessor` [configuration help][spanmetricsprocessor_help]             | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/processor/spanmetricsprocessor          |
| `spanprocessor`                                                                    | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/processor/spanprocessor                 |
| `tailsamplingprocessor` [configuration help][tailsamplingprocessor_help]           | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/processor/tailsamplingprocessor         |

[attributesprocessor_help]: https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/v0.41.0/processor/attributesprocessor
[groupbyattrsprocessor_help]: ./docs/Configuration.md#group-by-attributes-processor
[groupbytraceprocessor_help]: ./docs/Configuration.md#group-by-trace-processor
[metricstransformprocessor_help]: ./docs/Configuration.md#metrics-transform-processor
[resourcedetectionprocessor_help]: ./docs/Configuration.md#resource-detection-processor
[resourceprocessor_help]: https://github.com/open-telemetry/opentelemetry-collector-contrib/blob/v0.41.0/processor/resourceprocessor
[routingprocessor_help]: ./docs/Configuration.md#routing-processor
[spanmetricsprocessor_help]: ./docs/Configuration.md#span-metrics-processor
[tailsamplingprocessor_help]: ./docs/Configuration.md#tail-sampling-processor
[filterprocessor_help]: ./docs/Configuration.md#filter-processor

### Exporters

#### Sumo Logic supported exporters

| Name                                                             | Source                                                                                         |
|------------------------------------------------------------------|------------------------------------------------------------------------------------------------|
| `sumologicexporter` [configuration help][sumologicexporter_help] | https://github.com/SumoLogic/sumologic-otel-collector/tree/main/pkg/exporter/sumologicexporter |

[sumologicexporter_help]: ./docs/Configuration.md#sumo-logic-exporter

#### Upstream exporters

| Name                                                                     | Source                                                                                                        |
|--------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------|
| `carbonexporter` [configuration help][carbonexporter_help]               | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/exporter/carbonexporter        |
| `fileexporter` [configuration help][fileexporter_help]                   | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/exporter/fileexporter          |
| `kafkaexporter` [configuration help][kafkaexporter_help]                 | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/exporter/kafkaexporter         |
| `loadbalancingexporter` [configuration help][loadbalancingexporter_help] | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/exporter/loadbalancingexporter |
| `loggingexporter` [configuration help][loggingexporter_help]             | https://github.com/open-telemetry/opentelemetry-collector/tree/v0.41.0/exporter/loggingexporter               |
| `otlpexporter`                                                           | https://github.com/open-telemetry/opentelemetry-collector/tree/v0.41.0/exporter/otlpexporter                  |
| `otlphttpexporter`                                                       | https://github.com/open-telemetry/opentelemetry-collector/tree/v0.41.0/exporter/otlphttpexporter              |

[carbonexporter_help]: ./docs/Configuration.md#carbon-exporter
[fileexporter_help]: ./docs/Configuration.md#file-exporter
[kafkaexporter_help]: ./docs/Configuration.md#kafka-exporter
[loadbalancingexporter_help]: ./docs/Configuration.md#load-balancing-exporter
[loggingexporter_help]: ./docs/Configuration.md#logging-exporter

### Extensions

#### Sumo Logic supported extensions

| Name                                                               | Source                                                                                           |
|--------------------------------------------------------------------|--------------------------------------------------------------------------------------------------|
| `sumologicextension` [configuration help][sumologicextension_help] | https://github.com/SumoLogic/sumologic-otel-collector/tree/main/pkg/extension/sumologicextension |

[sumologicextension_help]: ./docs/Configuration.md#sumo-logic-extension

#### Upstream extensions

| Name                                             | Source                                                                                                            |
|--------------------------------------------------|-------------------------------------------------------------------------------------------------------------------|
| `ballastextension`                               | https://github.com/open-telemetry/opentelemetry-collector/tree/v0.41.0/extension/ballastextension                 |
| `bearertokenauthextension`                       | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/extension/bearertokenauthextension |
| `storage` [configuration help][filestorage_help] | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/extension/storage                  |
| `healthcheckextension`                           | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/extension/healthcheckextension     |
| `oidcauthextension`                              | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/extension/oidcauthextension        |
| `pprofextension`                                 | https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/v0.41.0/extension/pprofextension           |
| `zpagesextension`                                | https://github.com/open-telemetry/opentelemetry-collector/tree/v0.41.0/extension/zpagesextension                  |

[filestorage_help]: ./docs/Configuration.md#file-storage-extension

<!-- markdownlint-enable MD013 -->
<!-- markdownlint-enable MD034 -->

## Contributing

For contributing guidelines, see [CONTRIBUTING](./CONTRIBUTING.md).
