# Server-Sent Events (SSE) demo
This is a demo of how to send Server-Sent Events from the backend to the frontend.
It shows that the frontend does not have to poll the backend for information. It can
just subscribe to the EventSource and the web page will be immediately updated when
the backend sends to this SSE.

It also shows how to implement a go channel broadcaster.

## Testing
Start backend by ```make run```  
Start frontend by opening frontend/index.html in a browser.

## Limitation
- SSE can handle a maximum of 6 connections per browser session.
- SSE can only send strings, not images, videos or streams.


## Inspiration from
https://www.youtube.com/watch?v=MuyYQWeBTyU