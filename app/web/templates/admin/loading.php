<style>
.loader {
  border: 16px solid #f3f3f3;
  border-radius: 50%;
  border-top: 16px solid #3498db;
  width: 120px;
  height: 120px;
  -webkit-animation: spin 2s linear infinite;
  animation: spin 2s linear infinite;
  position: absolute;
  margin-top: 30%;
  margin-left: 30%;
}

@-webkit-keyframes spin {
  0% { -webkit-transform: rotate(0deg); }
  100% { -webkit-transform: rotate(360deg); }
}

@keyframes spin {
  0% { transform: rotate(0deg); }
  100% { transform: rotate(360deg); }
}
.loading-container
{
  height: 400px;
  position: relative;
}
</style>


<!-- Modal -->
<div class="modal" id="template-loading-modal" role="dialog">
  <div class="modal-dialog">
    <div class="modal-content">
        <div class="loader"></div>
    </div>
  </div>
</div>