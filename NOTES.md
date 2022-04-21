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

- The Receive Method is super overloaded. I'd love to see that abstracted so the component is implementing much more concrete interfaces. Hide the complexity in the orchestration layer.

- 