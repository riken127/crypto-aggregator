const createServer = require('./server');
const port = process.env.PORT || 4000;

(async () => {
    const app = await createServer();
    app.listen(port, () => {
        console.log(`GraphQL API running at ::${port}`);
    });
})();