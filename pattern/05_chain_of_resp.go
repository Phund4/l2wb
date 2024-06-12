package pattern;

import (
	"fmt"
)

// Patient struct 
type Patient struct {
    name              string
    registrationDone  bool
    doctorCheckUpDone bool
    medicineDone      bool
}

// Department interface
type Department interface {
	setNext(Department)
	execute(*Patient)
}

// Reception struct
type Reception struct {
	next Department
}

// Назначаем обработчику следующий обработчик
func (r *Reception) setNext(next Department) {
	r.next = next
}

// Действие на этапе конкретного обработчика
func (r *Reception) execute(p *Patient) {
	if p.registrationDone {
		fmt.Printf("Patient with name %s already has passed registration", p.name)
		r.next.execute(p);
		return;
	}
	p.registrationDone = true;
	fmt.Printf("Patient with name %s has passed registration", p.name)
	r.next.execute(p);
}

// Doctor struct
type Doctor struct {
	next Department
}

// Назначаем обработчику следующий обработчик
func (d *Doctor) setNext(next Department) {
	d.next = next;
}

// Действие на этапе конкретного обработчика
func (d *Doctor) execute(p *Patient) {
	if p.doctorCheckUpDone {
		d.next.execute(p);
		fmt.Printf("Patient with name %s already has passed doctor's appointment", p.name)
		return;
	}
	p.doctorCheckUpDone = true;
	fmt.Printf("Patient with name %s has passed doctor's appointment", p.name)
	d.next.execute(p);
}

// Medical struct
type Medical struct {
	next Department
}

// Назначаем обработчику следующий обработчик
func (m *Medical) setNext(next Department) {
	m.next = next;
}

// Действие на этапе конкретного обработчика
func (m *Medical) execute(p *Patient) {
	if p.medicineDone {
		m.next.execute(p);
		fmt.Printf("Patient with name %s already has passed medical", p.name)
		return;
	}
	p.medicineDone = true;
	fmt.Printf("Patient with name %s has passed medical", p.name)
	m.next.execute(p);
}

// ChainOfRespUserCode Пользовательский код использования паттерна цепочки обязанностей
func ChainOfRespUserCode() {
    medical := &Medical{}

    //Set next for doctor department
    doctor := &Doctor{}
    doctor.setNext(medical)

    //Set next for reception department
    reception := &Reception{}
    reception.setNext(doctor)

    patient := &Patient{name: "abc"}
    //Patient visiting
    reception.execute(patient)
}

/*
Паттерн цепочка обязанностей
Применимость: 1) Когда программа должна выполнять цепочку действий, но 
каких и каким способом она сама не знает
2) Когда необходимо, что обработчики запросов выполнялись в строгом порядке
Плюсы: 1) Уменьшает зависимость между клиентом и обработчиками
2) Реализует принципы открытости/закрытости и единства ответственности
Минусы: 1) Написанный запрос может быть не использован
Пример использования: Пользователь делает запрос на сервер, и в качестве 
промежуточных операций типа валидации, serDe, аутентификации используются
несколько обработчиков, которые следуют друг за другом. 
*/