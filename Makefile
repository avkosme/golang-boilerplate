start:
	cd compose && ansible-playbook -i hosts.yml configure.yml --limit local
	docker-compose up -d
stop:
	docker-compose stop
