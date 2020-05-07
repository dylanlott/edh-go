<template>
<div class="card" style="width: 18rem;">
  <img class="card-img-top" :src="image" :alt="name">
  <div class="card-body">
    <h5 class="card-title">{{ name }}</h5>
    <p class="card-text">{{ text }}</p>
    <a href="#" class="btn btn-primary">View card image</a>
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
    'colorIdentity',
    'power',
    'toughness',
    'text',
    'cardTypes',
    'supertypes',
    'subtypes',
    'types,',
    'image',
    'text',
  ]
}
</script  >
