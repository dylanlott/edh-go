<template>
<div class="card" style="width: 18rem;">
  <img class="card-img-top" :src="image" :alt="name">
  <div class="card-body">
    <h5 class="card-title">{{ name }}</h5>
    <p class="card-text">{{ text }}</p>
    <p class="card-text">{{ types }}</p>
    <p class="card-text">{{ supertypes }}</p>
    <p class="card-text">{{ subtypes }}</p>
    <p class="card-text">{{ convertedManaCost }}</p>
    <p class="card-text">{{ colorIdentity }}</p>
    <p class="card-text">{{ image }}</p>
    <hr>
    <p>Counters</p>
  </div>
</div>
</template>
<script>
export default {
  name: 'Card',
  props: [
    'id',
    'name',
    'convertedManaCost',
    'manaCost',
    'colorIdentity',
    'power',
    'toughness',
    'text',
    'types',
    'supertypes',
    'subtypes',
    'types,',
    'image',
  ]
}
</script  >
