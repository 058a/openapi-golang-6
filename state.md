```mermaid
stateDiagram-v2
    [*] --> Wait: Start
    Wait --> Do: Start processing
    Stop --> Do: Restart processing
    Do --> Stop: Pause processing
    Do --> [*]: Finish processing
    Stop --> [*]: Finish processing
```