import { ApolloClient } from 'apollo-client';
import { HttpLink } from 'apollo-link-http';
import { WebSocketLink } from 'apollo-link-ws';
import { split } from 'apollo-link';
import { getMainDefinition } from 'apollo-utilities';
import { InMemoryCache } from 'apollo-cache-inmemory';

const cache = new InMemoryCache({
    addTypename: false
})
const httpLink = new HttpLink({
    uri: process.env.VUE_APP_BASE_URL || "http://localhost:8080/graphql",
});
const wsLink = new WebSocketLink({
    uri: process.env.VUE_APP_WEBSOCKET_URL || "ws://localhost:8080/graphql",
    options: {
        reconnect: true,
    },
});
const link = split(
    ({ query }) => {
        const { kind, operation } = getMainDefinition(query);
        return kind === 'OperationDefinition' && operation === 'subscription';
    },
    wsLink,
    httpLink,
);
export default new ApolloClient({
    link: link,
    cache: cache,
});
