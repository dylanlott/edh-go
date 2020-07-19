<template>
  <div>
    <app-message v-for="message of messages"
                 :key="message.id"
                 :message="message">
    </app-message>
  </div>
</template>

<script>
import gql from 'graphql-tag';
import Message from '@/components/Message';

export default {
  components: {
    'app-message': Message,
  },
  data() {
    return {
      messages: [],
    };
  },
  apollo: {
    messages() {
      const user = this.$currentUser();
      return {
        query: gql`
          {
            messages {
              id
              user
              text
              createdAt
            }
          }
        `,
        subscribeToMore: {
          // the subscribeToMore document defines the subscription 
          // being called / listened to
          document: gql`
            subscription($user: String!) {
              messagePosted(user: $user) {
                id
                user
                text
                createdAt
              }
            }
          `,
          variables: () => ({ user: user }),
          updateQuery: (prev, { subscriptionData }) => {
            // check that there is data, return old if there's nothing new
            if (!subscriptionData.data) {
              return prev;
            }
            // pull the latest data from the subscription.
            // messagePosted lines up to the gql name of the subscription 
            const message = subscriptionData.data.messagePosted;

            // if the previous message is already in the old messages,
            // return just the previous messages. 
            if (prev.messages.find((m) => m.id === message.id)) {
              return prev;
            }

            // make new array with old and new combined and return it 
            return Object.assign({}, prev, {
              messages: [message, ...prev.messages],
            });
          },
        },
      };
    },
  },
};
</script>
