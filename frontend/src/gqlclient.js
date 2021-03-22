import { ApolloClient } from 'apollo-client';
import { HttpLink } from 'apollo-link-http';
import { WebSocketLink } from 'apollo-link-ws';
import { split } from 'apollo-link';
import { getMainDefinition } from 'apollo-utilities';
import { InMemoryCache } from 'apollo-cache-inmemory';

const cache = new InMemoryCache()
const httpLink = new HttpLink({
    uri: 'http://localhost:8080/graphql',
});
const wsLink = new WebSocketLink({
    uri: 'ws://localhost:8080/graphql',
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
