<header class="main-header">
    <a href="#" class="logo">
      <span class="logo-mini"><b class="logo-mini-pefix">A</b>LT</span>
      <span class="logo-lg"><b class="logo-lg-pefix">Admin</b>LTE</span>
    </a>
    <nav class="navbar navbar-static-top" role="navigation">
      <a href="#" class="sidebar-toggle" data-toggle="push-menu" role="button">
        <span class="sr-only">Toggle navigation</span>
      </a>
      <div class="navbar-custom-menu">
        <ul class="nav navbar-nav">
          
          <li>
            <a href="#" id="template-back-button">
              <i class="fa fa-arrow-left fa-1x" aria-hidden="true"></i>
            </a>
          </li>

          <li>
            <a href="#" id="template-refresh-button">
              <i class="fa fa-refresh fa-1x" aria-hidden="true"></i>
            </a>
          </li>
          <!-- User Account Menu -->
          <li class="dropdown user user-menu">
            <!-- Menu Toggle Button -->
            <a href="#" class="dropdown-toggle" data-toggle="dropdown">
            <i class="fa fa-user-circle fa-1x" aria-hidden="true" style="color:white;"></i>
              <!-- The user image in the navbar-->
              <!-- hidden-xs hides the username on small devices so only the image appears. -->
              <span class="hidden-xs">
                <?= isset( $this->session ) ? $this->session->userdata['identity'] : 'Your Username' ?>
                <i class="caret"></i>
              </span>
            </a>
            <ul class="dropdown-menu">
              <!-- The user image in the menu -->
              <li class="user-header">
                  <i class="fa fa-user-circle fa-5x" aria-hidden="true" style="color:white;"></i>
                <p>
                  <?= isset( $this->session ) ? $this->session->userdata['identity'] : 'Your Username' ?>
                  <small><?= isset( $this->session ) ? $this->session->userdata['email'] : 'Your Email' ?></small>
                </p>
              </li>
              <!-- Menu Body -->
              <!-- Menu Footer-->
              <li class="user-footer">

                  <a href="<?= site_url('/admin/logout') ?>" class="btn btn-danger btn-block" style="background-color: #dd4b39!important; color: white!important;">Sign out</a>

              </li>
            </ul>
          </li>
          
        </ul>
      </div>
    </nav>
  </header>