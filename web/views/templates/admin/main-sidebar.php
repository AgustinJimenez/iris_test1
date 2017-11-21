 <!-- Left side column. contains the logo and sidebar -->
 <aside class="main-sidebar">

    <!-- sidebar: style can be found in sidebar.less -->
    <section class="sidebar">

      <!-- Sidebar Menu -->
      <ul class="sidebar-menu" data-widget="tree">
        <li class="header">MODULES</li>
        <li class="treeview">
          <a href="#">
            <i class="fa fa-map-marker" aria-hidden="true" style="color:white;"></i> <span><?= "Traces" ?></span>
            <span class="pull-right-container">
              <i class="fa fa-angle-left pull-right"></i>
            </span>
          </a>
          <ul class="treeview-menu">
            <li class="active">
              <a href="<?= site_url('backend/traces') ?>">
                <i class="fa fa-map-o" aria-hidden="true"></i> 
                <?= "Traces List" ?>
              </a>
              
              </li>
          </ul>
        </li>

        <li class="header">SETTINGS</li>
        <li class="treeview">
          <a href="#">
            <i class="fa fa-user-circle" aria-hidden="true" style="color:white;"></i> <span><?= "Users" ?></span>
            <span class="pull-right-container">
              <i class="fa fa-angle-left pull-right"></i>
            </span>
          </a>
          <ul class="treeview-menu">
            <li class="active">
              <a href="<?= site_url('admin/users') ?>">
                <i class="fa fa-users" aria-hidden="true"></i> 
                <?= "Users List" ?>
              </a>
              <a href="<?= site_url('admin/create_user') ?>">
                <i class="fa fa-user-plus" aria-hidden="true"></i> 
                <?= "Create User" ?>
              </a>
              <a href="<?= site_url('admin/create_group') ?>">
                <i class="fa fa-users" aria-hidden="true"></i> 
                <?= "Create Group" ?>
              </a>
              
              </li>
          </ul>
        </li>
      </ul>
      <!-- /.sidebar-menu -->
    </section>
    <!-- /.sidebar -->
  </aside>