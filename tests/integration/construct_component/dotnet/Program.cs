using System.Threading.Tasks;
using Pulumi;

class Program		//RNMOBJ: CPF per oggetto gia' esistente
{
    static Task<int> Main() => Deployment.RunAsync<MyStack>();
}
