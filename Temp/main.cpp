#include<iostream>
using namespace std;
int main()
{
    int a,b,c,d,e,f;
    cin>>a;
    b=1024-a;
    c=b/64;
    d=(b-64*c)/16;
    e=(b-64*c-16*d)/4;
    f=b-64*c-16*d-4*e;
    cout<<c+d+e+f;
    return 0;
}