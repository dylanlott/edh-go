<template>
<div>
  <!-- Page Content -->
  <div class="shell landing">
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
                      <p class="subtitle">Sign up now and get early access when we launch the Beta!</p>
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
                        class="button subscribe-button is-success is-medium">
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
        <article class="tile is-child notification is-white box has-text-centered">
          <h3><i class="fa fa-mobile fa-4x"></i></h3>
          <p class="title">Mobile First</p>
          <p class="subtitle">Play anytime, anywhere, on any device with an internet connection and a browser.</p>
        </article>
      </div>
      <div class="column">
        <article class="tile is-child notification is-white box has-text-centered">
          <h3><i class="fa fa-arrows-alt fa-4x"></i></h3>
          <p class="title">Flexible</p>
          <p class="subtitle">Our gameplay engine allows for any format to be played, but we mostly focused on EDH.</p>
        </article>
      </div>
      <div class="column">
        <article class="tile is-child notification is-white box has-text-centered">
          <h3><i class="fa fa-star fa-4x"></i></h3>
          <p class="title">Simple</p>
          <p class="subtitle">We worked hard to keep the interface intuitive and easy to use on mobile <i>and</i> desktop.</p>
        </article>
      </div>
      <div class="column">
        <article class="tile is-child notification is-white box has-text-centered">
          <h3><i class="fa fa-tachometer-alt fa-4x"></i></h3>
          <p class="title">Fast</p>
          <p class="subtitle">The app is built from the ground up from scratch to be as small, fast, and reliable as possible.</p>
        </article>
      </div>
    </section>
    <section class="columns level screenshot">
      <div class="column">
        <article class="box levels">
          <h1 class="title is-2 level-item">
          EDH-Go is the first Magic: The Gathering board state tracker of its kind.
          </h1>
          <a target="_blank" href="https://i.imgur.com/1YUGzki.png">
            <img class="level-item screen" src="https://i.imgur.com/1YUGzki.png" alt="Magic: The Gathering Board State Tracker">
          </a>
          <p class="level-item what-is-edhgo">
            <h1 class="title is-2 has-text-centered">What is EDH-Go?</h1>
            <p class="subtitle has-text-centered">
              EDH-Go is a board state tracker that puts the power to control the game in the players hands.
              We don't try to handle all the logic of every single card interaction in the game.
              Instead, we leave that up to the players and focus on tracking those changes in the cleanest, most efficient way possible.
              This allows us to serve mobile users with a robust platform that's still performant even on limited internet connections.
            </p>
            <p class="level-item">
              <h1 class="title is-2 has-text-centered">The EDH Social Contract</h1>
              <p class="subtitle has-text-centered">
                We rely on the EDH Social Contract. This means we make a few pragmatic tradeoffs to achieve our end goal of performance on any platform.
                Player actions are auditable and visible to all other opponents, but their values are not. For example, drawing a card will be visible to
                all players, but what you drew is not visible. Anything that you'd be able to see a player do at the table, you'll be able to tell if a player
                did that in EDH-Go. 
              </p>
              <p class="subtitle has-text-centered">
                This works because we fundamentally believe in people's ability to self-monitor and be honest and transparent. If you want to cheat,
                everyone will see it and you're only playin' yourself.
                <br>

              </p>
              <h1 class="title is-1 has-text-centered">
                <b>Be a good player, be a good person.</b>
              </h1>
            </p>
          </p>
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
<style type="text/css" lang="scss" scoped>
@import "@/scss/custom.scss";

#mc_embed_signup{
  clear:left; 
  font:14px Helvetica,Arial,sans-serif; 
  width:100%;
}
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
.benefits {
  margin: .25rem;
}
.screenshot {
  margin: .25rem;
}
.screen {
  border: solid 2px $primary;
  border-radius: 5px;
}
.what-is-edhgo {
  margin: 1rem;
}
</style>
