import { createServer } from 'http';

const server = createServer((req, resp) => {
    resp.writeHead(200, { 'Content-Type': 'text/plain' });
    resp.end('Hello World\n');
});

const PORT = process.env.PORT || 9999

server.listen(PORT, () => {
    console.log(`Server running at http://localhost:${PORT}/`);
})