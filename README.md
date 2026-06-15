# 🚀 Projeto Korp - Desafio DevOps

Serviço HTTP em Golang com Docker, Nginx, Prometheus, Grafana e automação com Ansible.

## 📋 Sobre o projeto

Este projeto foi desenvolvido como solução para um desafio técnico de DevOps. O objetivo é criar um serviço HTTP em Golang, containerizado com Docker, com proxy reverso via Nginx, monitoramento com Prometheus/Grafana e automação completa com Ansible.

## 🛠️ Tecnologias utilizadas

| Tecnologia | Finalidade |
|------------|------------|
| **Go (Golang)** | Serviço HTTP |
| **Docker / Docker Compose** | Containerização |
| **Nginx** | Proxy reverso |
| **Prometheus** | Coleta de métricas |
| **Grafana** | Visualização de métricas |
| **Ansible** | Automação do ambiente |

## 📦 Pré-requisitos

- Linux (Fedora, Ubuntu ou outra distribuição)
- Git
- Ansible instalado

## 🚀 Como executar o projeto

```bash
# 1. Clone o repositório
git clone https://github.com/wallyson14/projeto-korp.git
cd projeto-korp

# 2. Execute o playbook do Ansible (um único comando)
sudo ansible-playbook ansible/playbook.yml
