#include <stdio.h>
#include <stdlib.h>

struct node
{
	int data;
	struct node *next;
};
struct node *first;

void create(int a[],int n)
{
	struct node *temp,*last;
	first=(struct node *)malloc(sizeof(struct node));
	first->data=a[0];
	first->next=NULL;
	last=first;
	for(int i=1;i<n;i++)
	{
		temp=(struct node *)malloc(sizeof(struct node));
		temp->data=a[i];
		temp->next=NULL;
		last->next=temp;
		last=temp;
	}
}

void display(struct node *p)
{
	while(p)
	{
		printf("%d ",p->data);
		p=p->next;
	}
}

void insertLast(struct node *p,int x)
{
	struct node *temp;
	temp->data=x;
	temp->next=NULL;
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//NOT COMPLIETE>>>>>>>>>>>>>>>>>>>>>>>>>>>
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
	//
}

	
		
int main()
{
	int ar[]={991,13,14,4,1};
	create(ar,5);
	display(first);
}
