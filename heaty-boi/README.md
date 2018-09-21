# heaty-boi

Experimental golang program used to keep your mining farm, and GPUs at the best temperature possible. The hotter your GPU, the harder and less efficient it works increasing your total power costs for both rig operation and facility cooling, while decreasing your hashing rate.

Only works for NVIDIA, AMD will *never* be supported, scram. We're strictly team green over here.
Presumably works on windows, but only ever tested on linux. Want windows? shoooo, *nix only shop here bud.

Heaty Boi will be rolled out in a few different phases:

* Basic temperature watching [x]
    * If a single GPU on the rig reaches a certain temperature stop all miners
    * Sleep for pre-determind amount of time
* More advanced temperature watching [ ]
    * If a sinlge GPU on the rig reaches a certain temperature stop all miners
    * Watch chip temps, once all chips reach a certain temperature start miners
* Big Boi temperature watching [ ]
    * Each GPU, or sets of GPUs will have its own mining process attached to it
    * If a GPU in a set reaches a certain temperature stop the miners associated with that set
    * Once that set reaches a certain temperature, resume miners