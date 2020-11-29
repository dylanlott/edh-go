<template>
    <div>Opponents</div>
</template>

<script>
import {
    boardstatesSubscription
} from '@/gqlQueries';

export default {
    name: 'Opponents',
    data () {
        return {
        }
    },
    props: {
        PlayerIDs: Array,
    },
    apollo: {
        opponents() {
            this.PlayerIDs.forEach((player, i) => {
                console.log('opponents#player: ', i, player)
                // setup Apollo query and return the results here. 
            })
        }
    }
}
</script>

<style>

</style>