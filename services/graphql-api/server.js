const express = require('express');
const { graphqlHTTP } = require('express-graphql');
const schema = require('./graphql/schema');
const root = require('./graphql/resolvers');
const { connect } = require('./redis');

/**
 * Creates and configures an Express server for the GraphQL API.
 * It connects to the Redis database, sets up the GraphQL endpoint,
 * and provides a health check endpoint.
 * @returns {Promise<express.Application>} - A promise that resolves to an Express application instance.
 */
async function createServer() {
    await connect();
    const app = express();
    app.use('/graphql', graphqlHTTP({
        schema,
        rootValue: root,
        graphiql: true
    }));

    app.get('/healthz', (req, res) => res.json({ status: 'ok' }));
    return app;
}

module.exports = createServer;