var builder = WebApplication.CreateBuilder(args);
var app = builder.Build();
app.MapGet("/", () => "Welcome to the C# Dashboard!");
Console.WriteLine("Updated message!");

app.Run();
