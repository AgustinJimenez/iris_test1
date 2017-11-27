<?php  
    if( !isset($message) )
        $message = '';

    $type = isset($message['type'])?$message['type']:' btn-github';
    $message = isset($message['body']) ? $message['body'] : ( !empty($message)?$message:'' );
    $message .= function_exists('validation_errors')?validation_errors():'';
    
?>

<?php if( !empty($message) ): ?>
    <div class="container-fluid">
        <div class="alert alert-<?= $type ?> fade in">
            <a href="#" class="close" data-dismiss="alert" aria-label="close">&times;</a>
            <?= $message ?>
        </div>
    </div>
<?php endif;?>
<style>
      label.error {
    color: red !important;
}
</style>