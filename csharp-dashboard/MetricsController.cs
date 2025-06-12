using Microsoft.AspNetCore.Mvc;
using System.Net.Http;
using System.Text.Json;
using System.Threading.Tasks;
using System.Threading;

[ApiController]
[Route("api/[controller]")]
public class MetricsController : ControllerBase
{
    private readonly HttpClient _httpClient;

    public MetricsController(IHttpClientFactory httpClientFactory)
    {
        _httpClient = httpClientFactory.CreateClient();
    }

    [HttpGet]
    public async Task<IActionResult> GetMetrics(CancellationToken cancellationToken)
    {
        var response = await _httpClient.GetAsync("http://python-reporter:5000/metrics", cancellationToken);
        if (!response.IsSuccessStatusCode)
        {
            return StatusCode((int)response.StatusCode, "Failed to get metrics");
        }

        var content = await response.Content.ReadAsStringAsync(cancellationToken);
        var data = JsonSerializer.Deserialize<object>(content);

        return Ok(data);
    }
}
