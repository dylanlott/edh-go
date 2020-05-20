<template>
<div>
  <!-- Page Content -->
  <div class="container">

    <!-- Jumbotron Header -->
    <header class="jumbotron my-4">
      <div class="row">
        <div class="col-12">
          <h1 class="display-3">Welcome to EDH Go!</h1>
          <h3>
            The mobile-first way to play Magic: The Gathering.
          </h3>
        </div>
      </div>
      <div class="row">
        <div class="col-sm-12 col">

          <!-- Begin Mailchimp Signup Form -->
          <link href="//cdn-images.mailchimp.com/embedcode/horizontal-slim-10_7.css" rel="stylesheet" type="text/css">
          <div id="mc_embed_signup">
          <form action="https://gmail.us18.list-manage.com/subscribe/post?u=2c6b66f6b94ee6dfbf9524986&amp;id=36cf9e0d6c" method="post" id="mc-embedded-subscribe-form" name="mc-embedded-subscribe-form" class="validate" target="_blank" novalidate>
              <div id="mc_embed_signup_scroll">
          	<label for="mce-EMAIL">
              <h5>Sign up now and get early access when we launch the Beta!</h5>
            </label>
            <br>
          	<input type="email" value="" name="EMAIL" class="form-control email" id="mce-EMAIL" placeholder="email address" required>
              <!-- real people should not fill this in and expect good things - do not remove this or risk form bot signups-->
              <div style="position: absolute; left: -5000px;" aria-hidden="true"><input type="text" name="b_2c6b66f6b94ee6dfbf9524986_36cf9e0d6c" tabindex="-1" value=""></div>
              <div class="clear"><input type="submit" value="Subscribe" name="subscribe" id="mc-embedded-subscribe" class="btn btn-primary btn-large btn-block signup-button"></div>
              </div>
          </form>
          </div>
          <!--End mc_embed_signup-->

        </div>
      </div>

      <!-- <a href="#" class="btn btn-primary btn-lg">Call to action!</a> -->
    </header>

    <!-- Page Features -->
    <div class="row text-center">

      <div class="col-lg-3 col-md-6 mb-4">
        <div class="card h-100">
          <!-- <img class="card-img-top" src="http://placehold.it/500x325" alt=""> -->
          <h3><i class="fa fa-mobile fa-5x"></i></h3>
          <div class="card-body">
            <h4 class="card-title">Mobile First</h4>
            <p class="card-text">Play anytime, anywhere, on any machine with an internet connection and a browser.</p>
          </div>
          <!-- <div class="card-footer">
            <a href="#" class="btn btn-primary">Find Out More!</a>
          </div> -->
        </div>
      </div>

      <div class="col-lg-3 col-md-6 mb-4">
        <div class="card h-100">
          <h3><i class="fa fa-gear fa-5x"></i></h3>
          <div class="card-body">
            <h4 class="card-title">Flexible</h4>
            <p class="card-text">Our gameplay engine allows for any format to be played, but we mostly focused on EDH.</p>
          </div>
          <!-- <div class="card-footer">
            <a href="#" class="btn btn-primary">Find Out More!</a>
          </div> -->
        </div>
      </div>

      <div class="col-lg-3 col-md-6 mb-4">
        <div class="card h-100">
          <h3><i class="fa fa-pencil fa-5x"></i></h3>
          <div class="card-body">
            <h4 class="card-title">Simple</h4>
            <p class="card-text">We worked hard to keep the interface intuitive and easy to use on mobile <i>and</i> desktop.</p>
          </div>
          <!-- <div class="card-footer">
            <a href="#" class="btn btn-primary">Find Out More!</a>
          </div> -->
        </div>
      </div>

      <div class="col-lg-3 col-md-6 mb-4">
        <div class="card h-100">
          <h3><i class="fa fa-dashboard fa-5x"></i></h3>
          <div class="card-body">
            <h4 class="card-title">Fast</h4>
            <p class="card-text">The app is built from the ground up from scratch to be as small and fast as possible, while still providing a good gameplay experience.</p>
          </div>
          <!-- <div class="card-footer">
            <a href="#" class="btn btn-primary">Find Out More!</a>
          </div> -->
        </div>
      </div>

    </div>
    <!-- /.row -->

  </div>
  <!-- /.container -->

  <!-- Footer -->
  <footer class="py-5 bg-dark">
    <div class="container">
      <p class="m-0 text-center text-white">Copyright &copy; Hivemind {{ year }}</p>
    </div>
    <!-- /.container -->
  </footer>
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
<style type="text/css">
  #mc_embed_signup{clear:left; font:14px Helvetica,Arial,sans-serif; width:100%;}
  /* Add your own Mailchimp form style overrides in your site stylesheet or in this style block.
     We recommend moving this block and the preceding CSS link to the HEAD of your HTML file. */
     .signup-button {
       margin: .75rem 0rem;
     }
</style>
