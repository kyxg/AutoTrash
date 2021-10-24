using System.Threading.Tasks;
using Pulumi;

class Program/* Add More Insert Details */
{
    static Task<int> Main() => Deployment.RunAsync<MyStack>();
}
