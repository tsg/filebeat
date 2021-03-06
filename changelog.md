Log to track "changes" from logstash-forward to filebeat

*



Questions:

 * .logstash-forwarder: Should we have a script that renames / rewrites the file to .filebeat on upgrade?
 * File comaprison for windows is not really implemented yet, see filecompare_windows.go
 * File size limit for config file was removed. In case this is required, it must be handled by libbeat
 * How well did the Windows integration work?
 * Config file: Type is required and moved out of fields. Fields can optional for additional information
 * Deadtime was used in the code but I couldn't find an example in any config? I added it now to the config
 * Config: input field was introduced with option "stdin" or "log". Default is "log". Idea is that in the future
   also full files could be read (fsriver) (InputBeat). Fifoin or streams could be also added. (see also https://github.com/elastic/logstash-forwarder/issues/525)
 ** All config files must end with .yml. In case a directory is passed as config path, all .yml files in this directory
     will be interpreted as config files and merged
 * Check exact deadtime behaviour and where it should be applied
 * Default config must be introduced again:

```
var defaultConfig = &struct {
   	fileDeadtime string
   }{
   	fileDeadtime: "24h",
   }
```

* Profiler option was removed as part of libbeat. Currently the profiler stopped after 60s. Should this be added to libbeat?
  Profiler options were also removed.
```
		go func() {
			time.Sleep(60 * time.Second)
			pprof.StopCPUProfile()
			panic("60-seconds of profiling is complete. Shutting down.")
		}()

		// Recover function for panic part

		defer func() {
			p := recover()
			if p == nil {
				return
			}

			fmt.Printf("recovered panic: %v", p)
			os.Exit(exitStat.faulted)
		}()
```


* Go trough error messages and check if the texts are good
* netTimeout was removed as config option as this is part of libbeat
* syslog config params (log-to-syslog and syslog) removed, as this is part of libbeat
* All command line options were also translated to config files options
* Getting config from env was removed. I think a better method like getting it from es should be used: https://github.com/elastic/logstash-forwarder/pull/435
* What should we do about multiple configs? Just provide some docs? https://github.com/elastic/logstash-forwarder/issues/136 currently working with -c for beat -config for dirs

Notes:
* Should every config entry have a name -> make it possible to know from which config entry something comes.
  Can it be theat config files overlap -> double indexing of a file?
* All beats should "namespace" the config file, otherwise would could have overlaps. Means also for packetbeat, everything should be under "packetbeat"
* We need general concept / code that command line args overwrite config options
* All command line options must be available as config options for the beats

Next with priority
* Multi line support
* Filtering support
