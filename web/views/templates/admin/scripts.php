
<?php foreach( isset($this->load->extra_js) ? $this->load->extra_js : []  as $file_path ):?>
    <script src="<?= $file_path ?>"></script>
<?php endforeach;?>

<?= $this->load->view('templates/admin/global-scripts/main', null, true); ?>
