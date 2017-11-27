<script src="<?= base_url('public/js/jquery.dataTables.min.js') ?>"></script>
<script src="<?= base_url('public/js/dataTables.bootstrap.min.js') ?>"></script>
<script src="<?= base_url('public/js/jquery.slimscroll.min.js') ?>"></script>

<script>
	$('.dataTable').DataTable
	({
		dom: "<'row'<'col-xs-12'<'col-xs-6'l><'col-xs-6'p>>r>"+
                "<'row'<'col-xs-12't>>"+
                "<'row'<'col-xs-12'<'col-xs-6'i><'col-xs-6'p>>>",
       "deferRender": true,
       "order": [[ 0, "desc" ]],
       //serverSide: true,
      'paging'      : true,
      "paginate": true,
      processing: true,
      'lengthChange': true,
      'responsive': true,
      'searching'   : true,
      'ordering'    : true,
      "filter": true,
      "sort": true,
      'info'        : true,
      'autoWidth'   : true,
      language: {
                    processing:     "Procesando...",
                    search:         "Buscar",
                    lengthMenu:     "Mostrar _MENU_ Elementos",
                    info:           "Mostrando de _START_ al _END_ registros de un total de _TOTAL_ registros",
                    infoEmpty:      "Mostrando 0 registros",
                    infoFiltered:   ".",
                    infoPostFix:    "",
                    loadingRecords: "Cargando Registros...",
                    zeroRecords:    "No existen registros disponibles",
                    emptyTable:     "No existen registros disponibles",
                    paginate: {
                        first:      "Primera",
                        previous:   "Anterior",
                        next:       "Siguiente",
                        last:       "Ultima"
                    }
                }
    })
</script>