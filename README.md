# easy-workflow
目前处于内测阶段，预计2月内发布正式版本

## easy-workflow是什么  
这是一个用纯GO语言开发的简单易用工作流引擎,可以集成到GO项目中，也可以作为Web Api Server单独运行

## 它有什么功能
除了基本的流程处理，它包含以下增强功能：  
1、可自定义节点节点事件。目前支持开始与结束事件。  
2、自由驳回功能：可以任意驳回到上游任何节点。  
3、直接提交到上一个驳回我的节点。    
### 功能2、3是基于国内特色情况考虑设计  
故事背景:假设A、B、C、D、E 共5个任务节点节点，E节点（老板）比较任性(向来如此)  
1、正常情况下,老板要打回一个流程的方式应该是驳回到D，D再驳回到C...最终打回到流程提交人A  
但是老板有钱任性，想要直接打回到A。此时，可以使用“自由驳回”功能，满足老板直接驳回，不留情面的需求    
2、“直接提交到上一个驳回我的节点”功能，是基于这么一种情况：  
A被领导E直接驳回，于是按照领导指示做修改后重新提交，传统情况下。流程需要重新流过B、C、D几个主管。  
一次两次还好，奈何老板各种不满意各种驳回，每次都要主管们再审核一遍。这样不仅效率低，还影响领导对员工的情绪。    
此时老板发话：芝麻绿豆大的事，B、C、D不用再参合了，小A你直接提我这边吧！  
此时使用"直接提交到上一个驳回我的节点"，A直接提交到上次驳回他的E，皆大欢喜。  
**简言之，这个流程引擎考虑国情，更好的满足老板多样化的需求**

