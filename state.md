```mermaid
stateDiagram-v2
    [*] --> Wait
    Wait --> Do
    Stop --> Do
    Do --> Stop
    Do --> [*]
    Stop --> [*]
```