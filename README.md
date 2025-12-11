# vcalendar
### **Goals of this project**:
- Implement a basic Terminal application that listens to the user and retrieves google calendar events
  - v1 only supports view all events on a given date 
    - v2 will try to implement creation, deletion and update of events
    - v2 will have more time specific features right now we can only list for the whole day but would like to add more features like listing events for a specific time range or filtering events based on their duration.


The core process transforms speech into calendar actions through several stages.

For a visual representation of this pipeline, [view the system data flow diagram](dataflow.png).

### **Process Stages**
1.  **Audio Input**: Captured via PortAudio.
2.  **Speech-to-Text**: Processed by Vosk over WebSocket.
3.  **Text Analysis**:
    *   Date extraction using the `when` library.
    *   Intent classification via Ollama embeddings and Qdrant.
4.  **Calendar Action**: Execution via Google Calendar API.
