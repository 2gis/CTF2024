<?php


class First{
    private $foo;
    public function foo(){
        return $this->foo->current();
    }
}

class Second{
    private $foo = false;

    public function foo(){
        if ($this->foo)
            return true;

    }
}

class Serializer
{
    public function __construct()
    {
        $this->first = new First();
        $this->second = new Second();
    }

    private $first;
    private $second;

    public function __wakeup()
    {
        if ($this->first->foo()) {
            if ($this->second->foo()) {
                echo '2GIS.CTF{byp4s5s_PHP_s3ri4l1z4t10n}';
            }
        }
    }
}

unserialize($_GET["payload"]);
?>
