<template>
<div>
  <!-- Page Content -->
  <div class="shell">
    <!-- Jumbotron Header -->
    <section class="hero is-medium is-primary is-bold welcome">
      <div class="hero-body">
        <div class="container">
          <h1 class="title is-1">
            Welcome to EDH-Go!
          </h1>
          <div class="subtitle">
              <!-- Begin Mailchimp Signup Form -->
              <link
                href="//cdn-images.mailchimp.com/embedcode/horizontal-slim-10_7.css"
                rel="stylesheet"
                type="text/css">
              <div id="mc_embed_signup">
                <form
                  action="https://gmail.us18.list-manage.com/subscribe/post?u=2c6b66f6b94ee6dfbf9524986&amp;id=36cf9e0d6c"
                  method="post"
                  id="mc-embedded-subscribe-form"
                  name="mc-embedded-subscribe-form"
                  class="validate"
                  target="_blank"
                  novalidate>
                  <div id="mc_embed_signup_scroll">
                  	<label for="mce-EMAIL">
                      <h5>Sign up now and get early access when we launch the Beta!</h5>
                    </label>
                    <br>
                  	<b-input
                      type="email"
                      value=""
                      name="EMAIL"
                      class="email"
                      id="mce-EMAIL"
                      placeholder="email address"
                      required/>
                      <!-- real people should not fill this in and expect good things - do not remove this or risk form bot signups-->
                      <div style="position: absolute; left: -5000px;" aria-hidden="true">
                        <input
                        type="text"
                        name="b_2c6b66f6b94ee6dfbf9524986_36cf9e0d6c"
                        tabindex="-1"
                        value="">
                      </div>
                      <div class="clear">
                        <input
                        type="submit"
                        value="Subscribe"
                        name="subscribe"
                        id="mc-embedded-subscribe"
                        class="button subscribe-button">
                      </div>
                    </div>
                  </form>
                </div>
              <!--End mc_embed_signup-->
          </div>
        </div>
      </div>
    </section>
    <!-- Benefits Tiles -->
    <section class="columns benefits">
      <div class="column">
        <article class="tile is-child notification is-white box">
          <h3><i class="fa fa-mobile fa-5x"></i></h3>
          <p class="title">Mobile First</p>
          <p class="subtitle">Play anytime, anywhere, on any device with an internet connection and a browser.</p>
        </article>
      </div>
      <div class="column">
        <article class="tile is-child notification is-white box">
          <h3><i class="fa fa-arrows-alt fa-5x"></i></h3>
          <p class="title">Flexible</p>
          <p class="subtitle">Our gameplay engine allows for any format to be played, but we mostly focused on EDH.</p>
        </article>
      </div>
      <div class="column">
        <article class="tile is-child notification is-white box">
          <h3><i class="fa fa-star fa-5x"></i></h3>
          <p class="title">Simple</p>
          <p class="subtitle">We worked hard to keep the interface intuitive and easy to use on mobile <i>and</i> desktop.</p>
        </article>
      </div>
      <div class="column">
        <article class="tile is-child notification is-white box">
          <h3><i class="fa fa-tachometer-alt fa-5x"></i></h3>
          <p class="title">Fast</p>
          <p class="subtitle">The app is built from the ground up from scratch to be as small, fast, and reliable as possible.</p>
        </article>
      </div>
    </section>
    <section class="columns level">
      <div class="column">
        <article class="box levels">
          <h3 class="title">
          EDH-Go is the first Magic: The Gathering board state tracker of its kind.
          </h3>
          <!-- TODO: Add image of the app here -->
          <img class="level-item" src="http://www.placehold.it/1500x900" alt="">
        </article>
      </div>
    </section>
    <!-- Footer -->
    <footer class="footer">
      <div class="content has-text-centered">
        <p><strong>EDH-Go</strong> by <a href="https://dylanlott.com/">Hivemind.</a></p>
      </div>
    </footer>
  </div>
</div>
</template>
<script type="text/javascript">
export default {
  name: 'landing',
  data () {
    return {
      year: 2020
    }
  }
}
</script>
<style type="text/css" scoped>
#mc_embed_signup{clear:left; font:14px Helvetica,Arial,sans-serif; width:100%;}
.signup-button {
  margin: .75rem 0rem;
}
.benefits {
  margin: 1rem -.7rem;
}
.subscribe-button {
  margin: 1rem 0rem;
}
.footer {
  width: 100%;
}
</style>
