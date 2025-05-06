
# Oficina Introdução à Linguagem Go

## Primeiros passos para configuração do ambiente local:

### Pré Requisitos:
 - Realizar a instalação de um editor de código ou IDE, as recomendações são o VSCode ou Goland.
<div style="text-align: center;">  
     <img src="/assets/vscode.png" style="margin-right: 100px" width="100"/>                       
     <img src="/assets/goland.png" style="margin-right: 100px" width="100"/>
</div>

## 1. Baixar o Instalador
 - Acesse o site oficial do Go para baixar a versão mais recente: [Download Go](https://go.dev/doc/install)
 - Escolha a versão correspondente ao seu sistema operacional:
   1. Windows: .msi
   2. Linux: .tar.gz
   3. macOS: .pkg
## 2. Instalar o Go

Como sempre vou deixar um link de um vídeo para auxiliar caso surja alguma dúvida.
 - [Tutorial de instalação Go.](https://youtu.be/eJq_D9at6ec?si=NQeV1cZcozKjjgVC)

### Windows

1. Execute o arquivo `.msi` baixado.
2. Siga as instruções do assistente de instalação.
3. Verifique se o Go foi instalado corretamente abrindo o **Terminal** e executando:
   ```sh
   go version
   
   go version go1.24.1  👈 o output esperado é esse
   ```

### Linux

1. Extraia o arquivo baixado para `/usr/local`:
   ```sh
   sudo tar -C /usr/local -xzf go<versao>.linux-amd64.tar.gz
   ```
2. Adicione o caminho do Go ao `PATH` editando `~/.bashrc` ou `~/.profile`:
   ```sh
   echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bashrc
   source ~/.bashrc
   ```
3. Verifique a instalação:
   ```sh
   go version
   
   go version go1.24.1  👈 o output esperado é esse
   ```

### macOS

1. Execute o arquivo `.pkg` baixado.
2. Siga as instruções do instalador.
3. Verifique a instalação abrindo o **Terminal** e executando:
   ```sh
   go version
   
   go version go1.24.1  👈 o output esperado é esse
   ```

### 4. Por fim, faça o `git clone` deste repositório na sua máquina para acompanhar e realizar os exercícios durante a oficina!
### 5. Tá pronto o sorvetinho! Nos vemos novamente no dia da oficina  🐈🚀

### Caso surja alguma dúvida, entre em contato por e-mail (liprog@ufcspa.edu.br) ou pelo instagram (@liprog.ufcspa).
