
<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <title>AdminLTE 2 | Starter</title>
  <meta content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no" name="viewport">
  <?php foreach( isset($this->load->extra_css) ? $this->load->extra_css : [] as $file_path ):?>
    <link rel="stylesheet" href="<?= $file_path ?>"/>
  <?php endforeach;?>
</head>