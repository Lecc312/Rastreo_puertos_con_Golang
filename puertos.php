<?php
    ini_set('max_execution_time', 0);
    $ip = "caja-negra.com.mx";
    $max=65535;
    $min=0;
    $array=array();
    echo ("Analizando: ".$ip."\n");
    $inicio=new DateTime(date('h:i:s'));
    for ($puerto = $min; $puerto <= $max; $puerto++){
        echo ($puerto."/".$max."\n");
        if ($fp=@fsockopen($ip,$puerto,$ERROR_NO,$ERROR_STR,(float)0.3))
        {
            fclose($fp);
            array_push($array,"Puerto: ".$puerto.": ONLINE \n");
        }
    }
    $fin=new DateTime(date('h:i:s'));
    $tiempo = $inicio->diff($fin);
    $total=count($array);
    echo("---------------------------------------------------------------------------------------------------------------- \n");
    echo ("Analisis de ".$ip." completo realizado en ".$tiempo->format('%H horas %i minutos %s segundos')." \n");
    sleep(2);
    for ($i=0;$i<$total;$i++){
        echo($array[$i]."\n");
        usleep(200000); //espera de 0.2 segundos
    }    
    echo("Fin del analisis \n");
?>
