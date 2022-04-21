- Just trying to get a working config took a bit. Eventually went with 

nodes:
  - name: generator
    outputs:
    - fake
    metric_generator:
      spawn_interval: 10s
  - name: fake
    fake_metric_remote_write:
      credential: {}

- Was a little hard to find exactly which components I could use. I found The fake metric receiver in the code, but wasn't sure how to invoke it in yaml. Then I found the monster Node struct. Yuck.

- Modified the fake remote write to print metrics so I could iterate quickly. Not sure if better way.

- The Receive Method is super overloaded. I'd love to see that abstracted so the component is implementing much more concrete interfaces. Hide the complexity in the orchestration layer.

- Step 1. Add a new relabeler type to config. Panic immediately: `panic: unable to handle node named relabel`

- Took me a while to find the if/else block of doom in orchestrator.go. Some self/dynamic registration of components would go a long way.

- "You have not implemented the actor interface" - lots of copy/paste and trial/error to get compiling again

- I think I sent a bad message type to my output (single metric vs array), and I got no feedback what was wrong, just that no metrics were flowing there.

- exchange.Metric has no way to add or change a label without copying whole metric. This makes it a little annoying to do simple relabeling. I see an attempt to make everything immutable, but I worry about the memory allocations for every step of the way.

