<!DOCTYPE html>
<html>
<?= $this->load->view('templates/admin/head', null, true); ?>
<!--
|---------------------------------------------------------|
| SKINS         | skin-blue                               |
|               | skin-black                              |
|               | skin-purple                             |
|               | skin-yellow                             |
|               | skin-red                                |
|               | skin-green                              |
|---------------------------------------------------------|
|LAYOUT OPTIONS | fixed                                   |
|               | layout-boxed                            |
|               | layout-top-nav                          |
|               | sidebar-collapse                        |
|               | sidebar-mini                            |
|---------------------------------------------------------|
-->
<body class="hold-transition skin-blue sidebar-mini" style="background-color: #222d32;">
<div class="wrapper" >
    <?= $this->load->view('templates/admin/main-header', null, true); ?>
    <?= $this->load->view('templates/admin/main-sidebar', null, true); ?>
    <?= $this->load->view('templates/admin/loading', null, true); ?>
    <div class="content-wrapper">
     
        <?= $this->load->view('templates/admin/message', null, true); ?>

        <section class="content container-fluid">
            <div class="row">
                <div class="col-xs-12">
