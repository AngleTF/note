```linux
#!/bin/bash

#mysqlVar
host="127.0.0.1"
port=3306
user="you name"
password="you password"

#filePath
rootBackupPath="/data/backup"
date=`date "+%Y-%m-%d"`
mysqldumpPath="/data/mysql/bin/mysqldump"
logPath="/data/backup/backup.log"
backupPath="${rootBackupPath}/${date}"

sqlSuffix=".sql"
compressSuffix=".tar.bz2"

#backup of databases
dbs=(dbname1 dbname2)

#create path
mkdir -p ${rootBackupPath} ${backupPath}

echo -e "\n[ "`date "+%Y-%m-%d %H:%M:%S"`" ]\nbackup to directory ${backupPath}" >> ${logPath}


#backup new sql
backupSql(){
	
	for dbname in $@
	do

	echo "dump ${dbname} to ${backupPath}/${dbname}${sqlSuffix}" >> ${logPath}
	${mysqldumpPath} -u${user} -p${password} -h${host} -P${port} -R -E ${dbname} > "${backupPath}/${dbname}${sqlSuffix}" 2>> ${logPath}
	
	done
}

#remove old sql
rmOlderBackup(){
	echo -e "remove ${rootBackupPath} older backup\n\n" >> ${logPath}
	#find ${rootBackupPath} -type d -a -name "[0-9][0-9][0-9][0-9]-[0-9][0-9]-[0-9][0-9]" -a -mtime +7 -exec rm -r {} \;
	find ${rootBackupPath} -regextype "posix-egrep" -regex "^${rootBackupPath}/[0-9]{4}-[0-9]{2}-[0-9]{2}(${compressSuffix}){0,1}$" -a -mtime +7 -exec rm -rf {} \;  >> ${logPath}
}

#bz2 compress
compress(){
	echo "compress ${backupPath} to ${backupPath}${compressSuffix}" >> ${logPath}
	tar -Pjcvf ${backupPath}${compressSuffix} ${backupPath} 2>> ${logPath}
}


backupSql ${dbs[*]}

compress

rmOlderBackup


```