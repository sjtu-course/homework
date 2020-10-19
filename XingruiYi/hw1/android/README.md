## Homework 1 易兴睿 Android

1、采用`SharedPreferences`来实现房间名与用户名的存取。具体通过`editor.setText`与`editor.putString`来进行存取操作。

2、在写入Preference后，通过`editor.apply`而非`editor.commit`实现数据写回，原因在于`editor.apply`通过异步的方式将数据写回硬盘，而`editor.commit`则是当即写入，在数据量较大的时候会造成app暂时的卡顿。



