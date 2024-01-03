package pokeapi
import (
	"io"
	"encoding/json"
	"net/http"
	"fmt"
)
func (c *Client) ListLocationArea(pageUrl *string) (locationAreaResp, error){
   endpoint:="/location-area"
   fullUrl:=baseUrl+endpoint

   if pageUrl != nil{
	fullUrl = *pageUrl
   }

 // check the catch

 dat, ok :=c.cache.Get(fullUrl)
if ok {
	// cache hit
	fmt.Println("cache hit")
	locationAreaResps:=locationAreaResp{}
err:=json.Unmarshal(dat, &locationAreaResps)

if err!=nil {
	return locationAreaResp{}, err
}

return locationAreaResps, nil
}

fmt.Println("cache miss")

req, err:=http.NewRequest("GET", fullUrl, nil)
if err!=nil{
	return locationAreaResp{}, err
}
resp, err :=c.httpClient.Do(req)

if err != nil {
	return locationAreaResp{}, err
}
defer resp.Body.Close()

if resp.StatusCode>399{
	return locationAreaResp{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
}
dat,err=io.ReadAll(resp.Body)

if err!=nil {
	return locationAreaResp{}, err
}

locationAreaResps:=locationAreaResp{}
err=json.Unmarshal(dat, &locationAreaResps)

if err!=nil {
	return locationAreaResp{}, err
}
c.cache.Add(fullUrl, dat)
return locationAreaResps, nil

}


func (c *Client) GetLocationArea(locationAreaName string) (LocationArea, error){
	endpoint:="/location-area/" + locationAreaName
	fullUrl:=baseUrl+endpoint
 
	
 
  // check the catch
 
  dat, ok :=c.cache.Get(fullUrl)
 if ok {
	 // cache hit
	 fmt.Println("cache hit")
	 locationArea:=LocationArea{}
 err:=json.Unmarshal(dat, &locationArea)
 
 if err!=nil {
	 return LocationArea{}, err
 }
 
 return locationArea, nil
 }
 
 fmt.Println("cache miss")
 
 req, err:=http.NewRequest("GET", fullUrl, nil)
 if err!=nil{
	 return LocationArea{}, err
 }
 resp, err :=c.httpClient.Do(req)
 
 if err != nil {
	 return LocationArea{}, err
 }
 defer resp.Body.Close()
 
 if resp.StatusCode>399{
	 return LocationArea{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
 }
 dat,err=io.ReadAll(resp.Body)
 
 if err!=nil {
	 return LocationArea{}, err
 }
 
 locationArea:=LocationArea{}
 err=json.Unmarshal(dat, &locationArea)
 
 if err!=nil {
	 return LocationArea{}, err
 }
 c.cache.Add(fullUrl, dat)
 return locationArea, nil
 
 }