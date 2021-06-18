<template>
<div>
    <h3>{{ boardstate.User.Username }}</h3>
    <div class="battlefield">
        <div v-for="(card, i) in boardstate.Field" :key="i">
            <Card v-bind="card" />
        </div>
    </div>

    <!-- <div class="exiled">
        <div v-for="(card, i) in boardstate.Exiled" :key="i">
            <Card v-bind="card" />
        </div>
    </div> -->

    <div class="revealed">
        <div v-for="(card, i) in boardstate.Revealed" :key="i">
            <Card v-bind="card" />
        </div>
    </div>

    <div class="library">
        <div v-for="(card, i) in boardstate.Library" :key="i">
            <!-- TODO -->
            <!-- <Card v-bind="card"/> -->
        </div>
    </div>
</div>
</template>

<script>
import Card from '@/components/Card';
import { mapState } from 'vuex';
export default {
  components: { Card },
    name: 'Opponent',
    data () {
        return {
        }
    },
    computed: mapState({
        boardstates: state => state.BoardStates.boardstates,
    }),
    props: {
        boardstate: Object
    },
}
</script>

<style>

</style>