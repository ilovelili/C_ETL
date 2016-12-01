# CimpressJP ETL Framework

## Introduction 
The main construct in etl framework is Pipeline. A Pipeline has a series of PipelineStages, which will each perform some type of data processing, and then send new data on to the next stage. Each PipelineStage consists of one or more DataProcessors, which are responsible for receiving, processing, and then sending data on to the next stage of processing. DataProcessors each run in their own goroutine, and therefore all data processing can be executing concurrently.

Here is a conceptual drawing of a fairly simple Pipeline:
<pre>
+--Pipeline------------------------------------------------------------------------------------------+
|                                                                       PipelineStage 3              |
|                                                                      +---------------------------+ |
|  PipelineStage 1                 PipelineStage 2          +-JSON---> |  CSVWriter                | |
| +------------------+           +-----------------------+  |          +---------------------------+ |
| |  SQLReader       +-JSON----> | Custom DataProcessor  +--+                                        |
| +------------------+           +-----------------------+  |          +---------------------------+ |
|                                                           +-JSON---> |  SQLWriter                | |
|                                                                      +---------------------------+ |
+----------------------------------------------------------------------------------------------------+
</pre>

## Contact
mju@cimpress.com